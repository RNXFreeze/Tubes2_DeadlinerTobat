export default function TreePage({algorithmType, searchElement, setSearchElement}) {
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
                    disabled={searchElement.trim() === ' '}
                    className={`px-4 py-2 ${
                        searchElement.trim() === ''
                            ? 'bg-purple-300 text-white cursor-not-allowed'
                            : 'bg-purple-500 text-white hover:bg-purple-400'
                    }`}
                    onClick={() => {
                        console.log("Search clicked:", setSearchElement, algorithmType);
                    }}>
                    Search
                </button>
            </div>
        </div>
      </div>

      {/* Placeholder area for recipe tree */}
      <div className="min-w-[800px] min-h-[500px] border rounded bg-white p-4">
        <p className="text-gray-500">Visualisasi recipe tree akan tampil di sini!</p>
      </div>
    </main>
  );
}