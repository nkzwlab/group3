import React from 'react';
import './App.css';

class App extends React.Component {
    render = () => (
        <div className="App">
            <KadaiList />
        </div>
  );
}


class KadaiList extends React.Component {
    constructor(props) {
        super(props);
        fetch('http://localhost:8080/kadai?user_id=1').then(response => response.json()).then(this.setKadais).catch(error => console.error(error));
        this.state = {kadais: []};
    }

    setKadais = data => {
        console.log('fetched: ', data);
        this.setState({kadais: data});
    }

    render = () => {
        console.log("state kadais: ", this.state.kadais);
        const kadaiItems = this.state.kadais.map(kadai =>
            <KadaiItem
                key={kadai.id}
                kadai={kadai} />
        );
        return kadaiItems;
    }
}

class KadaiItem extends React.Component {

    render = () => {
        const kadai = this.props.kadai;

        return (
            <div className="kadaiItem">
                <p>title: {kadai.title}</p>
                <p>content: {kadai.content}</p>
                <p>draft: {kadai.draft}</p>
            </div>
        );
    }
}

export default App;