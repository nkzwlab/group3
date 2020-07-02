import React from 'react';
import './App.css';

class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            loggedIn: false
        };
    }

    setUser = user => {
        console.log("logged in as :", user.login_name)

        if (!user.login_name || !user.id) {
            throw user["error"];
        }

        this.setState({loggedIn: true, user: {
            loginName: user.login_name,
            id: user.id
        }});
    }

    login = loginName => {
        const url = `http://localhost:8080/user?login_name=${loginName}`
        fetch(url)
            .then(response => response.json())
            .then(this.setUser)
            .catch(error => {
                let text = error;
                if (error["error"]) {
                    const text = error.error;
                }
                const newUrl =  `http://localhost:8080/user/new?login_name=${loginName}`;
                return fetch(newUrl, {
                    method:"POST",
                    headers: new Headers({
                        'Accept': 'application/json',
                    })
                })
            })
            .then(response => response.json())
            .then(this.setUser)
            .catch(error => {
                let text = error;
                if (error["error"]) {
                    const text = error.error;
                }
                alert("エラーが発生しました：" + text);
            });
    }
            body: body
    render = () => {
        const node = (this.state.loggedIn)
            ? (
                <div className="mainContent">
                    <KadaiList user={this.state.user} />
                </div>)
            : <Login onLogin={this.login} onSubmit={this.login} />;


        return (
            <div className="App">
                <h1 className="pageHeader">課題管理サービス</h1>
                {node}
            </div>
        );
    }
}



class Login extends React.Component {
    constructor(props) {
        super(props);
        this.state = {loginName: ""};
    }

    handleChange = e => {
        const loginName = e.target.value;
        this.setState({loginName});
    }

    handleSubmit = e => {
        this.props.onSubmit(this.state.loginName);
        e.preventDefault();
    }

    render = () => (
        <form className="loginForm" onSubmit={this.handleSubmit}>
            <input type="text" value={this.state.loginName} onChange={this.handleChange} />
            <input type="submit" value="ログイン" />
        </form>
    );
}




class KadaiList extends React.Component {
    constructor(props) {
        super(props);
        this.refreshKadais()

        this.state = {
            kadais: [],
        };
    }

    refreshKadais = () => {
        console.log("refresh");
        const kadaiURL = `http://localhost:8080/kadai?user_id=${this.props.user.id}`;
        fetch(kadaiURL)
            .then(response => response.json())
            .then(this.setKadais)
            .catch(error => {
                let text = error;
                if (error["error"]) {
                    const text = error.error;
                }
                alert("エラーが発生しました：" + text);
            });
    }

    setKadais = data => {
        if (data !== null && data !== undefined) {
            const kadais = data.map(kadai => {
                return {
                    id: kadai.id,
                    userId: kadai.user_id,
                    title: kadai.title,
                    content: kadai.content,
                    draft: kadai.draft,
                    editing: false
                };
            })
            this.setState({kadais: kadais});
        } else {
            this.setState({kadais: []});
        }
    }

    editKadai = kadai => {
        const kadais = this.state.kadais.map(v => {
            if (v.id === kadai.id) {
                v.editing = true;
            }
            return v;
        });
        this.setState(kadais);
    }

    updateDone = id => {
        const doneURL = `http://localhost:8080/kadai/done?kadai_id=${id}`;

        const body = new URLSearchParams();
        body.append('kadai_id', id);

        fetch(doneURL, {
            method: "POST",
        })
            .then(response => alert("done!"))
            .catch(err => {
                console.error("error!: ", err);
            })
            .finally(() => {
                this.refreshKadais();
            })
    }

    render = () => {
        const kadaiItems = this.state.kadais.map(kadai => (
            <KadaiItem
                user={this.props.user}
                key={kadai.id}
                kadai={kadai}
                onDone={this.updateDone}
                refresh={this.refreshKadais}
                editing={kadai.editing}
            />)
        );

        return (
            <div>
                <h2>未提出課題一覧</h2>
                <button onClick={this.refreshKadais}>再読み込み</button>
                {kadaiItems}
                <h2>新しい課題を登録する</h2>
                <PostKadai refresh={this.refreshKadais}
                           user={this.props.user} />
            </div>
        )
    }
}

class KadaiItem extends React.Component {
    constructor(props) {
        super(props);
        this.state = {kadai: this.props.kadai};
    }

    startEdit = e => {
        console.log("edit started");
        let kadai = this.state.kadai;
        kadai.editing = true;
        this.setState({kadai});
    }

    finishEdit = kadai => {
        console.log("edit finished");
        kadai.editing = false;
        this.setState({kadai})
    }

    handleDone = e => {
        this.props.onDone(this.state.kadai.id);
        this.props.refresh();
    }

    render = () => {
        const kadai = this.state.kadai;

        const node = (kadai.editing)
            ? <UpdateKadai kadai={kadai}
                           finishEdit={this.finishEdit}
                           refresh={this.props.refresh} />
            : <ShowKadai kadai={kadai}
                         handleEdit={this.startEdit}
                         handleDone={this.handleDone}/>;

        return (
            <div className="kadaiItem">
                {node}
            </div>
        );
    }
}

class ShowKadai extends React.Component {
    constructor(props) {
        super(props);
        this.kadai = this.props.kadai;
    }

    render = () => {
        return (
            <div>
                <h3 className="kadaiTitle">{this.kadai.title}</h3>
                <ul>
                    <li>課題内容:
                        <p className="kadaiBody">{this.kadai.content}</p>
                    </li>
                    <li> 下書き:
                        <p className="kadaiBody">{this.kadai.draft}</p>
                    </li>
                </ul>
                <button className="edit" onClick={this.props.handleEdit} >編集</button>
                <button className="done" onClick={this.props.handleDone} >提出完了</button>
            </div>
        );
    };
}

class KadaiEditor extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            kadai: this.props.kadai
        };
    }

    handleSubmit = e => {
        this.props.onSubmit(this.state.kadai);
        this.setState({kadai: {
            title: "",
            content: "",
            draft: ""
        }})
        e.preventDefault();
    }

    handleChange = (key, value) => {
        let kadai = this.state.kadai;
        kadai[key] = value;
        this.setState({kadai});
    }

    render = () => {
        const kadai = this.state.kadai;
        return (<div>
            <form onSubmit={this.handleSubmit}>
                <ul>
                    <KadaiForm type="title" value={kadai.title} changeHandler={this.handleChange} />
                    <KadaiForm type="content" value={kadai.content} changeHandler={this.handleChange} />
                    <KadaiForm type="draft" value={kadai.draft} changeHandler={this.handleChange} />
                    <input type="submit" value="完了" />
                </ul>
            </form>
        </div>)
    }
}

class UpdateKadai extends React.Component {
    constructor(props) {
        super(props);
        this.kadai = this.props.kadai;
    }

    updateKadai = kadai => {
        const baseURL = `http://localhost:8080/kadai/update?`;

        const body = new URLSearchParams();
        body.set('kadai_id', kadai.id);
        body.set('title', kadai.title);
        body.set('content', kadai.content);
        body.set('draft', kadai.draft);

        fetch(baseURL + body.toString(), {
            method: "POST",
            headers: new Headers({
                'Accept': 'application/json',
            }),
            body: body
        })
            .then(response => response.json())
            .then(data => {
                if (data["error"]) {
                    alert(data.error);
                    return
                }
                console.log(data)
            })
            .catch(error => {
                let text = error;
                if (error["error"]) {
                    const text = error.error;
                }
                alert("エラーが発生しました：" + text);
            })
            .finally(() => {
                this.props.finishEdit(kadai);
                console.log("end");
                this.props.refresh();
            });

    }

    render = () => (
        <KadaiEditor kadai={this.kadai} onSubmit={this.updateKadai} />
    )
}


class PostKadai extends React.Component {
    constructor(props) {
        super(props);
        this.kadai = {
            user_id: this.props.user.id,
            title: "",
            content: "",
            draft: "",
            editing: true
        };
        this.user = this.props.user;
    }

    postNewKadai = kadai => {
        const postURL = "http://localhost:8080/kadai/new";

        const body = new URLSearchParams();
        body.append('user_id', this.user.id);
        body.append('title', kadai.title);
        body.append('content', kadai.content);
        body.append('draft', kadai.draft);

        fetch(postURL, {
            method: "POST",
            headers: new Headers({
                'Accept': 'application/json',
            }),
            body: body
        }).then(response => response.json())
          .then(data => {
            if (data["error"]) {
                console.error(data.error);
            }
            console.log(data)
            this.props.refresh();
        }).catch(error => {
                let text = error;
                if (error["error"]) {
                    const text = error.error;
                }
                alert("エラーが発生しました：" + text);
            });

    }

    render = () => (
        <KadaiEditor kadai={this.kadai} onSubmit={this.postNewKadai} />
    );
}

class KadaiForm extends React.Component {
    handleChange = e => {
        this.props.changeHandler(this.props.type, e.target.value);
        e.preventDefault();
    }
    render = () => {
        const type = this.props.type;

        const showText = {
            "title": "課題名",
            "content": "課題内容",
            "draft": "下書き"
        };
        const text = showText[type];

        const input = (type === "title")
            ? <input
                type="text"
                value={this.props.value}
                onChange={this.handleChange}
                className="kadaiEditTitle" />
            : <textarea
                type="text"
                value={this.props.value}
                onChange={this.handleChange}
                className="kadaiEditBody" />;

        return (
            <li>
                <p>{text}</p>
                {input}
            </li>
        )
    }
}

export default App;
