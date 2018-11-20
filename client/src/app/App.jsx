import React from 'react';
import axios from 'axios'

class App extends React.Component {

  state = {
    name: '',
    email: '',
    password: ''
  }

  onChange = type => e => this.setState({ [type]: e.target.value })

  submit = (e) => {
    e.preventDefault();
    axios.post('/api/user', this.state)
  }

  render() {
    return(
      <div style={{ display: 'flex' }}>
        name
        <input onChange={this.onChange('name')} />
        email
        <input onChange={this.onChange('email')} />
        password
        <input onChange={this.onChange('password')} />

        <button onClick={this.submit}>
          SUBMIT
        </button>
      </div>
    );
  }
}

export default App;
