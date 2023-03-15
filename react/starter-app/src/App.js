import logo from './logo.svg';
import './App.css';
import React, { Component } from 'react';
import Table from "./Table";

class App extends Component {
  state = {
    characters: [

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
  }

  removeCharacter = (index) => {
    const { characters } = this.state

    this.setState({
      characters: characters.filter((character, i) => {
        return i !== index
      }),
    })
  }

  render() {
    const { characters } = this.state

    return (
      <div className='container'>
        <Table charactersData={characters} removeCharacter={this.removeCharacter} />
      </div>
    )
  }
}


export default App
