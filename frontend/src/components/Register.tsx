import React, {SyntheticEvent, useState} from 'react';

const Register = (props: { onSubmit: any}) => {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [confirmPassword, setConfirmPassword] = useState('');
    const [message, setMessage] = useState('');

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();

        if (password !== confirmPassword)
            setMessage("Le nouveau mot de passe ne correspond pas");
        else {
            const response = await fetch('https://localhost:8000/api/register', {
                method: 'POST',
                headers: {'Content-Type': 'application/json'},
                credentials: 'include',
                body: JSON.stringify({
                    name,
                    email,
                    password
                })
            });

            const content = await response.json();
            if (content.id !== undefined){
                props.onSubmit();
            } else {
                setMessage(content.message);
            }
        }
    }

    return (
        <form onSubmit={submit}>
            <div className="form-group">
                <label htmlFor="name">Nom d'utilisateur</label><br />
            <input className="form-control" name="name" placeholder="Nom d'utilisateur" required
                   onChange={e => setName(e.target.value)}
            />
            </div><br />

            <div className="form-group">
                <label htmlFor="email">Adresse courriel</label><br />
            <input type="email" className="form-control" name="email" placeholder="Courriel" required
                   onChange={e => setEmail(e.target.value)}
            />
            </div><br />

            <div className="form-group">
                <label htmlFor="password">Mot de passe</label><br />
            <input type="password" className="form-control" name="password" placeholder="Mot de passe" required
                   onChange={e => setPassword(e.target.value)}
            />
            </div><br />

            <div className="form-group">
                <label htmlFor="confirmPassword">Confirmer nouveau mot de passe</label><br />
                <input type="password" className="form-control" name='confirmPassword' placeholder="Confirmer nouveau mot de passe" required
                   onChange={e => setConfirmPassword(e.target.value)}
                />
            </div><br />

            <button className="w-100 btn btn-lg btn-primary" type="submit">Soumettre</button>
            <p style={{color: 'red'}}>{message}</p>
        </form>
    );
};

export default Register;