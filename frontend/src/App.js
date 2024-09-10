
import './App.css';
import React, { useState, useEffect } from 'react';

function App() {
  const [persons, setPersons] = useState([]);
  const [firstName, setFirstName] = useState('');
  const [lastName, setLastName] = useState('');

  const fetchPersons = async () => {
    try {
      const response = await fetch('http://localhost:3000/api/persons')
      const data = await response.json();
      console.log("Persons from API:", data); 
      setPersons(data);
    } catch (error) {
      console.error('Error fetching persons:', error);
    }
  }

  useEffect(() => {
    console.log("Fetching persons...");
    fetchPersons();
  }, []);
  
  console.log("Rendering persons:", persons);

  const handleSubmit = async () => {
    const newPerson = { firstname: firstName, lastname: lastName};

    try {
      const response = await fetch('http://localhost:3000/api/person', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(newPerson),
      });

      if (response.ok) {
        fetchPersons();
        setFirstName('');
        setLastName('');
      } else {
        console.error('Error adding person');
      }
    } catch (error) {
      console.error('Error connecting to API:', error);
    }
  }

  return (
    <div className="middle">
      <h1>TestRestApi</h1>
      <div className="input-container">
        <input type="text" placeholder="Firstname" id="firstName" value={firstName} onChange={(e) => setFirstName(e.target.value)}>
        </input>
        <input type="text" placeholder="Lastname" id="lastName" value={lastName} onChange={(e) => setLastName(e.target.value)}>
        </input>
        <button onClick={handleSubmit}>Add person</button>
      </div>

      <div className="list">
        <ul id="person-list">
          {persons && persons.length > 0 ? (
            persons.map((person, index) => (
              <li key={index}>
                {person.firstname} {person.lastname}
              </li>
            ))
          ) : (
              <li>No persons available</li> // Alternativ text om listan är tom
          )}
        </ul>
      </div>
    </div>
  );
}

export default App;
