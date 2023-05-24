import {SyntheticEvent, useState} from 'react';
import { Spinner } from 'react-bootstrap';
import {Redirect} from "react-router-dom";

const Login = (props:any) => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
   
    const [blockMessage] = useState('');
    const [isForgetPasswordModalOpen, setIsForgetPasswordModalOpen] = useState(false);
    const [code2FA, setcode2FA] = useState('');
    const [loading, setLoading] = useState(false);

    const submit = async (e: SyntheticEvent) => {
        setLoading(true);
        e.preventDefault();
        await props.onLogin(email, password)
        setLoading(false);
    }

    const submit2fa = async (e: SyntheticEvent) => {
        setLoading(true);        
        e.preventDefault();
        await props.on2FA(code2FA)
        setLoading(false);
    }

    if (props.redirect) {
        return <Redirect to="/"/>;
    }

    if (props.show2FA) {
        return (
            <div className="form-signin">
                <form onSubmit={submit2fa}>
                    <h1 className="h3 mb-3 fw-normal">Veuillez entrer le code</h1>
                    <h2 className="h3 mb-3 fw-normal">Position x: {props.posx}</h2>
                    <h2 className="h3 mb-3 fw-normal">Position y: {props.posy}</h2>
                    <input type="code" className="form-control" placeholder="Code 2FA" required
                        onChange={e => setcode2FA(e.target.value)}
                    />
                    <p style={{color: 'red'}}>{props.message}</p>
                    <button className="w-100 btn btn-lg btn-primary" type="submit">Envoyer</button><br /><br />
                </form>
                {loading && <Spinner/>}
            </div>
        );
    }

    return (
        <div className="form-signin">
            {isForgetPasswordModalOpen && <ForgetPasswordModal isForgetPasswordModalOpen={isForgetPasswordModalOpen} setIsForgetPasswordModalOpen={setIsForgetPasswordModalOpen} />}
            <form onSubmit={submit}>
                <h1 className="h3 mb-3 fw-normal">Veuillez vous connecter</h1>
                <input type="email" className="form-control" placeholder="Courriel" required
                    onChange={e => setEmail(e.target.value)}
                />

                <input type="password" className="form-control" placeholder="Mot de passe" required
                    onChange={e => setPassword(e.target.value)}
                />

                <button className="w-100 btn btn-lg btn-primary" type="submit">Se connecter</button><br /><br />
                <button className="btn btn-link" onClick={() => setIsForgetPasswordModalOpen(true)}>Mot de passe oublié</button><br />
                <p style={{color: 'red'}}>{props.message}</p>
                <p style={{color: 'red'}}>{blockMessage}</p>
                {loading && <Spinner/>}
            </form>
        </div>
    );
};

function ForgetPasswordModal(props: any) {
    const [blockEmail, setBlockEmail] = useState('');
    const [blockMessage, setBlockMessage] = useState('');

    const askBlock = async (e: SyntheticEvent) => {
        e.preventDefault();
        const response = await fetch('https://localhost:8000/api/block', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                email: blockEmail
            }),
        });
        const content = await response.json();

        setBlockMessage(content.message);
    }
     
    return (
        <div className="modal">
        <div className="modal-content">
            <h2 >Mot de passe oublié</h2>
            <p>En cliquant sur <b>Confirmer</b>, vous consentez à bloquer votre compte temporairement afin qu'un administrateur vous assigne un mot de passe temporaire qui vous sera envoyé par courriel.</p>
            <form onSubmit={askBlock}>
                <div className="form-group">
                    <label htmlFor="blockEmail">Entrez votre adresse courriel</label><br />
                    <input type="email" className="form-control" name='blockEmail' placeholder="Courriel" required
                        onChange={e => setBlockEmail(e.target.value)}
                    />
                </div><br />

                <button className="w-100 btn btn-lg btn-primary" type="submit">Confirmer</button>
            </form>
            <button className="w-100 btn btn-lg btn-danger" onClick={() => props.setIsForgetPasswordModalOpen(false)} style={{marginTop: "20px"}}>Fermer</button>
            <p style={{color: 'red'}}>{blockMessage}</p>
        </div>
        </div>
    );
}
export default Login;