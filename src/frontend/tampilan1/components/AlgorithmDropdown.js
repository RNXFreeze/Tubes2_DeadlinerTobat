export default function AlgorithmDropdown({algorithmType, setAlgorithm, element, setElement, handleSearch}) {
    return(
        <div className="flex items-center gap-2">
            <div className="flex items-center border border-purple-300 rounded-lg overflow-hidden bg-white w-80">
                <input
                    type="text"
                    value={element}
                    onChange={(e) => setElement(e.target.value)}
                    placeholder="What element's recipe do you want to know?"
                    className="flex-grow px-4 py-2 outline-none"
                />
                <button 
                    disabled={element.trim() === ' '}
                    className={`px-4 py-2 ${
                        element.trim() === ''
                          ? 'bg-purple-300 text-white cursor-not-allowed'
                          : 'bg-purple-500 text-white hover:bg-purple-400'
                      }`}
                    onClick={handleSearch}
                >
                    Search
                </button>
            </div>
                <select
                    value={algorithmType}
                    onChange={(e) => setAlgorithm(e.target.value)}
                    className={"bg-purple-300 text-purple-900 font-bold border border-purple-300 px-4 py-2 rounded-lg hover:bg-white"}
                >
                    <option value="BFS">
                      BFS
                    </option>
                    <option value="DFS">
                      DFS
                    </option>
                    <option value="Bidirectional">
                      Bidirectional
                    </option>
                </select>
        </div>
    )
}