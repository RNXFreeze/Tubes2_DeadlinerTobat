'use client';

import { useState } from 'react';
import TreeVisualizer from './TreeVisualizer';

export default function TreePage({algorithmType, searchElement, setSearchElement}) {
  const [isLoading, setIsLoading] = useState(false);

  // DUMMY FUNCTION
  const searchRecipe = async (searchElement, algorithmType) => {
      // TODO: implement BFS/DFS here
      console.log('Searching for:', searchElement, 'using', algorithmType);
      await new Promise(resolve => setTimeout(resolve, 1000)); // simulasi loading
  };

  return (
    <main className="flex-grow p-6 overflow-auto">
      <div className="mb-6 text-center items-center">
        <h2 className="text-3xl font-bold mb-4">
            Little Alchemy Recipe Finder
        </h2>
        <div className="flex justify-center">
            <div className="flex items-center border border-purple-300 rounded-lg overflow-hidden bg-white w-80">
                <input
                    type="text"
                    value={searchElement}
                    onChange={(e) => setSearchElement(e.target.value)}
                    placeholder="What element's recipe do you want to know?"
                    className="flex-grow px-4 py-2 outline-none"
                />
                <button 
                    disabled={searchElement.trim() === '' || isLoading}
                    className={`px-4 py-2 ${
                        searchElement.trim() === '' || isLoading
                            ? 'bg-purple-300 text-white cursor-not-allowed'
                            : 'bg-purple-500 text-white hover:bg-purple-400'
                    }`}
                    onClick={async () => {
                        setIsLoading(true);
                        try {
                            await searchRecipe(searchElement, algorithmType);
                        } finally {
                            setIsLoading(false);
                        }
                    }}
                >
                    {isLoading ? (
                      <div className="flex items-center space-x-2">
                          <svg className="animate-spin h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                              <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                              <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v4l3-3-3-3v4a8 8 0 00-8 8z"></path>
                          </svg>
                          <span>Load</span>
                      </div>
                  ) : (
                      'Search'
                  )}
                </button>
            </div>
        </div>
      </div>

      {/* Placeholder area for recipe tree */}
      <div className="min-w-[800px] min-h-[500px] border rounded bg-white p-4">
        <TreeVisualizer />
      </div>
    </main>
  );
}