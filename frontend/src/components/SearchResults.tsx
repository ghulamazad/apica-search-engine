import React from 'react';
import { Record } from '../types/record';

interface SearchResultsProps {
  results: Record[];
  count: number;
  duration: string;
}

const SearchResults: React.FC<SearchResultsProps> = ({ results, count, duration }) => {
  return (
    <div className="mt-6 w-full">
      <p className="text-gray-700 mb-4">
        Found {count} results in {duration}.
      </p>

      <div className="w-full overflow-x-auto bg-white rounded-lg shadow">
        <table className="min-w-full table-auto text-sm text-left text-gray-700 border-collapse">
          <thead className="bg-gray-100">
            <tr>
              <th className="px-4 py-2 border border-gray-300">Message</th>
              <th className="px-4 py-2 border border-gray-300">MessageRaw</th>
              <th className="px-4 py-2 border border-gray-300">StructuredData</th>
              <th className="px-4 py-2 border border-gray-300">Tag</th>
              <th className="px-4 py-2 border border-gray-300">Sender</th>
              <th className="px-4 py-2 border border-gray-300">Groupings</th>
              <th className="px-4 py-2 border border-gray-300">Event</th>
              <th className="px-4 py-2 border border-gray-300">EventId</th>
              <th className="px-4 py-2 border border-gray-300">NanoTimeStamp</th>
              <th className="px-4 py-2 border border-gray-300">Namespace</th>
            </tr>
          </thead>
          <tbody>
            {results.map((record, index) => (
              <tr key={index} className="hover:bg-gray-50">
                <td className="px-4 py-2 border border-gray-300 max-w-[200px] truncate" title={record.Message}>
                  {record.Message}
                </td>
                <td className="px-4 py-2 border border-gray-300 max-w-[300px] truncate" title={record.MessageRaw}>
                  {record.MessageRaw}
                </td>
                <td className="px-4 py-2 border border-gray-300 max-w-[250px] truncate" title={record.StructuredData}>
                  {record.StructuredData}
                </td>
                <td className="px-4 py-2 border border-gray-300 max-w-[120px] truncate" title={record.Tag}>
                  {record.Tag}
                </td>
                <td className="px-4 py-2 border border-gray-300 max-w-[150px] truncate" title={record.Sender}>
                  {record.Sender}
                </td>
                <td className="px-4 py-2 border border-gray-300 max-w-[150px] truncate" title={record.Groupings}>
                  {record.Groupings}
                </td>
                <td className="px-4 py-2 border border-gray-300 max-w-[150px] truncate" title={record.Event}>
                  {record.Event}
                </td>
                <td className="px-4 py-2 border border-gray-300 max-w-[150px] truncate" title={record.EventId}>
                  {record.EventId}
                </td>
                <td className="px-4 py-2 border border-gray-300 max-w-[150px] truncate" title={record.NanoTimeStamp}>
                  {record.NanoTimeStamp}
                </td>
                <td className="px-4 py-2 border border-gray-300 max-w-[150px] truncate" title={record.Namespace}>
                  {record.Namespace}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default SearchResults;
