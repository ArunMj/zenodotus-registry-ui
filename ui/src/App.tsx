import * as React from 'react';
import './App.css';

import { RepositoryList } from './RepositoryList';

class App extends React.Component {
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <h1 className="App-title">Zenodotus registry ui</h1>
        </header>
          <RepositoryList/>
      </div>
    );
  }
}

export default App;
