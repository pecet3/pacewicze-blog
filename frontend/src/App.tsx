import './App.css'
import BasicTile from "./components/BasicTile.tsx";
import PostList from "./components/PostList.tsx";

function App() {
  
  return (
      <>
    <div className="flex justify-center items-center flex-col">
      <div  className="text-5xl font-bold ">Pacewicze blog
      </div>
     {/*<BasicTile title="Test" />*/}
      <PostList />
    </div>
      </>
  )
}

export default App
