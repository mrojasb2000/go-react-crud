function App() {
 
  return (
    <>
      <div>Hello World</div>
      <button onClick={async () => {
        const response = await fetch('/users')
        const data = await response.json()
        console.log(data)
      }}>Get users</button>
      <button onClick={async () => {
      const response = await fetch('http://localhost:3000/users', {
        method: 'POST',
        headers: {
          'Accept': 'application/json, text/plain, */*',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({name: "John Doe"})
        }).then(response => response.json())
        .then(response => console.log(response));
        console.log(response);
      }}>Create User</button>
    </>
  )
}

export default App