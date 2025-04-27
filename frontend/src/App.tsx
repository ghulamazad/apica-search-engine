import React, { useState } from 'react';
import axios from 'axios';
import SearchBar from './components/SearchBar';
import SearchResults from './components/SearchResults';
import { Record } from './types/record';
import UploadButton from './components/UploadButton';

const App: React.FC = () => {
  const [results, setResults] = useState<Record[]>([]);
  const [count, setCount] = useState(0);
  const [duration, setDuration] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const handleSearch = async (query: string) => {
    setLoading(true);
    setError('');
    try {
      const res = await axios.get('http://localhost:8080/search', { params: { q: query } });
      setResults(res.data.results);
      setCount(res.data.count);
      setDuration(res.data.duration);
    } catch (err) {
      console.error('Search error:', err);
      setError('Failed to search. Please try again.');
      setResults([]);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="flex flex-col min-h-screen bg-gray-100">
      <div className="p-4 bg-blue-600 text-white text-center text-2xl font-bold">
        Apica Search Engine
      </div>

      <div className="p-4 bg-white shadow-md">
        <SearchBar onSearch={handleSearch} />
        <UploadButton />
      </div>

      <div className="flex-grow overflow-auto p-4">
        {loading && (
          <div className="flex justify-center items-center h-full">
            <div className="loader ease-linear rounded-full border-8 border-t-8 border-gray-200 h-12 w-12"></div>
          </div>
        )}

        {!loading && error && (
          <p className="text-red-500 text-center my-4">{error}</p>
        )}

        {!loading && !error && !results  && (
          <p className="text-gray-500 text-center my-4">No results found.</p>
        )}

        {!loading && results && results.length > 0 && (
          <SearchResults results={results} count={count} duration={duration} />
        )}
      </div>
    </div>
  );
};

export default App;
