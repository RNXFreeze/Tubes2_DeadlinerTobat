import Navbar from '../components/navbar';

export default function Home() {
  return (
    <div className="bg-purple-100 min-h-screen">
      <Navbar />
      <main className="flex flex-col items-center justify-center p-6 gap-6 text-center">
        <h1 className="text-3xl font-bold text-center">Little Alchemy Recipe Finder</h1>
      <div className="inline-flex rounded-lg shadow-sm" role="group">
        <button className='bg-purple-300 text-purple-700 font-bold border border-purple-300 px-4 py-2 rounded-lg hover:bg-purple-200'>
          Shortest Recipe
        </button>
        <button className='bg-purple-300 text-purple-700 font-bold border border-purple-300 px-4 py-2 rounded-lg hover:bg-purple-200'>
          Multiple Recipe
        </button>
      </div>
      <div className="flex items-center gap-2">
        <div className="flex items-center border border-purple-300 rounded-lg overflow-hidden bg-white w-80">
          <input
            type="text"
            placeholder="What element's recipe do you want to know?"
            className="flex-grow px-4 py-2 outline-none"
          />
          <button className="bg-purple-500 text-white px-4 py-2 hover:bg-purple-400">
            Search
          </button>
        </div>
          <select className="bg-purple-300 text-purple-700 font-bold border border-purple-300 px-4 py-2 rounded-lg hover:bg-purple-200">
            <option value="">
              BFS
            </option>
            <option value="">
              DFS
            </option>
          </select>
      </div>
      <div className="text-center">
        <h2 className="font-bold mb-2">
          How many recipe do you want to see?
        </h2>
        <div className="flex items-center gap-2 justify center">
          <span>1</span>
          <input type="range" min="1" max='20' className="w-48" />
          <span>20</span>
        </div>
      </div>
      </main>
    </div>
  );
}