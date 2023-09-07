function App() {
  return (
    <>
      <div>Hello World</div>
      <button onClick={async () => {
        const response = await fetch('http://localhost:3000/users')
        const data = await response.json()
        console.log(data)
      }}>Get users</button>
    </>
  )
}

export default App