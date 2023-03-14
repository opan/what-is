import logo from './logo.svg';
import './App.css';
import React, { Component } from 'react';
import Table from "./Table";

class App extends Component {
  // render() {
  //   return (
  //     <div className="App">
  //       <header className="App-header">
  //         <img src={logo} className="App-logo" alt="logo" />
  //         <p>
  //           Edit <code>src/App.js</code> and save to reload.
  //         </p>
  //         <a
  //           className="App-link"
  //           href="https://reactjs.org"
  //           target="_blank"
  //           rel="noopener noreferrer"
  //         >
  //           Learn React
  //           Sudah berubah?
  //         </a>
  //       </header>
  //     </div>
  //   )
  // }
  render() {
    const characters = [
      {
        name: 'Charlie',
        job: 'Janitor',
      },
      {
        name: 'Mac',
        job: 'Bouncer',
      },
      {
        name: 'Dee',
        job: 'Aspring actress',
      },
      {
        name: 'Dennis',
        job: 'Bartender',
      },
      {
        name: 'Opan',
        job: 'Systems Engineer'
      },
    ]
    return (
      <div className='container'>
        <Table charactersData={characters} />
      </div>
    )
  }
}

export default App
