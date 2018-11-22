import React from 'react';
import axios from 'axios'

class App extends React.Component {

  state = {
    name: '',
    email: '',
    password: '',
    id: ''
  }

  onChange = type => e => this.setState({ [type]: e.target.value })

  submit = (e) => {
    e.preventDefault();
    axios.post('/api/user/', this.state)
  }

  getUser = e => {
    axios.get(`/api/user/${this.state.id}/`)
      .then(res => console.log(res))
  }

  render() {
    return(
      <div style={{ display: 'flex' }}>
        name
        <input value={this.state.name} onChange={this.onChange('name')} />
        email
        <input value={this.state.email} onChange={this.onChange('email')} />
        password
        <input value={this.state.password} onChange={this.onChange('password')} />

        <button onClick={this.submit}>
          SUBMIT
        </button>


        <input value={this.state.id} onChange={this.onChange('id')}>

        </input>

        <button onClick={this.getUser}>
          GET
        </button>
      </div>
    );
  }
}

export default App;
