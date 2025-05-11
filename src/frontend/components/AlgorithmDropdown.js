export default function AlgorithmDropdown({algorithmType, setAlgorithm}) {
    return(
        <div className="flex justify-center text-center gap-2">
            <select
                value={algorithmType}
                onChange={(e) => setAlgorithm(e.target.value)}
                className={"bg-purple-300 text-purple-900 font-bold border border-purple-300 px-4 py-2 rounded-lg hover:bg-purple-100"}
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