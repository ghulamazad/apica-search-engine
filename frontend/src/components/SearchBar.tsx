import React, { useState } from 'react';

interface SearchBarProps {
  onSearch: (query: string) => void;
}

const SearchBar: React.FC<SearchBarProps> = ({ onSearch }) => {
  const [query, setQuery] = useState('');

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (query.trim() !== '') {
      onSearch(query);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="flex w-full max-w-5xl mx-auto gap-4">
      <input
        type="text"
        value={query}
        placeholder="Search logs, events, messages..."
        onChange={(e) => setQuery(e.target.value)}
        className="flex-grow p-4 border border-gray-300 rounded-md text-lg focus:outline-none focus:ring-2 focus:ring-blue-400"
      />
      <button
        type="submit"
        className="bg-blue-600 hover:bg-blue-700 active:bg-blue-800 text-white font-semibold text-lg py-3 px-8 rounded-md shadow-md transition duration-300 ease-in-out transform hover:-translate-y-1 active:translate-y-0"
      >
        Search
      </button>
    </form>
  );
};

export default SearchBar;
