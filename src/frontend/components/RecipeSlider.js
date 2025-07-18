export default function RecipeSlider({ value, setValue }) {
  return (
    <div className="flex flex-col text-center gap-2">
      <h2 className="text-black font-bold">How many recipe do you want to see?</h2>
      <h3 className="text-black font"> (0 for all recipes, max shown = 50)</h3>
      <div className="flex items-center gap-2 justify-center">
        <span>0</span>
        <input
          type="range"
          min="0"
          max="20"
          value={value}
          onChange={(e) => setValue(e.target.value)}
          className="w-48"
        />
        <span>20</span>
      </div>
      <p className="text-sm text-purple-700">Selected: {value}</p>
    </div>
  );
}
