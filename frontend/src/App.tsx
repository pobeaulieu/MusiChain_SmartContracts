import React, { useEffect, useState } from 'react';
import './App.css';
import { BrowserRouter, Route } from 'react-router-dom';
import Nav from './components/Nav';
import Home from './pages/Home';
import Login from './pages/Login'
import Administration from './pages/Administration';
import Business from './pages/Business';
import Residential from './pages/Residential';
import Users from './pages/Users';

function App() {
    const [loggedUser, setLoggedUser] = useState({});
    const [message, setMessage] = useState("");
    const [redirect, setRedirect] = useState(false);
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [show2FA, setshow2FA] = useState(false);
    const [posx, setposx] = useState("");
    const [posy, setposy] = useState("");

    useEffect(() => {
        (
            async () => {
                const response = await fetch('https://localhost:8000/api/user', {
                    headers: {'Content-Type': 'application/json'},
                    credentials: 'include',
                });

                const content = await response.json();
                
                setLoggedUser(content);
            }
        )();
        // eslint-disable-next-line
    }, []);
    
    const onLogin = async (email : string, password : string) => {
        const loginResponse = await fetch('https://localhost:8000/api/login', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            credentials: 'include',
            body: JSON.stringify({
                email,
                password
            })
        });

        const contentLogin = await loginResponse.json();

        if (contentLogin.message === "success" || contentLogin.message === "Login success. Password need to be changed"){
            const policy = await fetch('https://localhost:8000/api/getloginpolicy', {
                headers: {'Content-Type': 'application/json'},
                credentials: 'include',
            });
            const data = await policy.json();
    
            if (data.twofa){
                setEmail(email)
                setPassword(password)

                const resp = await fetch('https://localhost:8000/api/getcode', {
                    headers: {'Content-Type': 'application/json'},
                    credentials: 'include',
                });
                const code = await resp.json()

                setposx(code.positionx)
                setposy(code.positiony)
                setMessage("")
                setshow2FA(true)
            }
            else{
                const response = await fetch('https://localhost:8000/api/user', {
                    headers: {'Content-Type': 'application/json'},
                    credentials: 'include',
                });
                const content = await response.json();
                setLoggedUser(content);
                setRedirect(true)
            }
        }
        
        else {
            setMessage(contentLogin.message)
        }
    }

    const on2FA = async (code: string) => {
        setMessage("");

        const response = await fetch('https://localhost:8000/api/login2fa', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            credentials: 'include',
            body: JSON.stringify({
                "email": email,
                "password": password,
                "code": code,
                "positionx" : posx.toString(),
                "positiony": posy.toString()
            })
        });

        const contentLogin = await response.json();

        if (contentLogin.message === "success"){

            setLoggedUser(contentLogin.user);
            setRedirect(true)
        }

        else{
            setMessage(contentLogin.error);
            
        }
    }

    const onLogout = async () => {
        setLoggedUser({});
        await fetch('https://localhost:8000/api/logout', {
            method: 'POST',
                headers: {'Content-Type': 'application/json'},
                credentials: 'include',
            });
        setRedirect(false)
        setshow2FA(false)

    }

    const loadUser =  async () => {
        setLoggedUser({});
        const response = await fetch('https://localhost:8000/api/user', {
            headers: {'Content-Type': 'application/json'},
            credentials: 'include',
        });
    
        const content = await response.json();
        
        setLoggedUser(content);
    }

    return (
        <div className="App">
            <BrowserRouter>
                <Nav loggedUser={loggedUser} onLogout={onLogout}/>

                <main>
                    <Route path="/" exact component={() => <Home loggedUser={loggedUser} loadUser={loadUser} message={message}/>}/>
                    <Route path="/login" exact component={() => <Login onLogin={onLogin} on2FA={on2FA} show2FA={show2FA} redirect={redirect} message={message} posx={posx} posy={posy}/>}/>
                    <Route path="/administration" exact component={() => <Administration loggedUser={loggedUser}/>}/>
                    <Route path="/clients/business" exact component={() => <Business loggedUser={loggedUser}/>}/>
                    <Route path="/clients/residential" exact component={() => <Residential loggedUser={loggedUser}/>}/>
                    <Route path="/users" exact component={() => <Users loggedUser={loggedUser}/>}/>
                </main>
            </BrowserRouter>
        </div>
    );
}

export default App;
