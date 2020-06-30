import React from 'react';
import './App.css';

function App() {
  return (
    <div className="App">
        <KadaiList />
    </div>
  );
}


class KadaiList extends React.Component {
    constructor(props) {
        super(props);
        this.kadais = [
            {
                id: 0,
                title: "jn_lecture",
                content: "グループでWebサービスを作る",
                draft: "課題管理サービスを開発中"
            }
        ];
    }

    render = () => {
        const kadaiItems = this.kadais.map(kadai =>
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
