import { createSignal } from 'solid-js'
import './App.css'

function App() {
  const [count, setCount] = createSignal(0)

  return (
    <>
      <h1>さう〜な</h1>    
      <button onClick={() => setCount((count) => count + 1)}>
        カウンター {count()}
      </button>
    </>
  )
}

export default App
