import './App.css'
import BasicTile from "./components/BasicTile.tsx";


function App() {
  
  return (
      <>
    <div className="flex justify-center items-center flex-col">
      <div  className="text-5xl font-bold ">Hello World!</div>
     <BasicTile title="Test" />
    </div>
      </>
  )
}

export default App
