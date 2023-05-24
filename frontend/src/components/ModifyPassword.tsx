import React, {SyntheticEvent, useState} from 'react';
import { Spinner } from 'react-bootstrap';
import { checkLogin } from './checkLogin';

const ModifyPassword = (props: { onSubmit: any, userEmail: string, isAdminRequest: boolean}) => {
    const [currentPassword, setCurrentPassword] = useState('');
    const [newPassword, setNewPassword] = useState('');
    const [confirmNewPassword, setConfirmNewPassword] = useState('');
    const [message, setMessage] = useState('');
    const [loading, setLoading] = useState(false);


    const submit = async (e: SyntheticEvent) => {
        setLoading(true);
        e.preventDefault();

        async function apiCallModifyPassword() {
            const response = await fetch('https://localhost:8000/api/modifypassword', {
                method: 'POST',
                headers: {'Content-Type': 'application/json'},
                credentials: 'include',
                body: JSON.stringify({
                    email: props.userEmail,
                    password: newPassword
                })
            });

            const content = await response.json();
            if (content.message === "success" || content.message === "Login success. Password need to be changed"){
                props.onSubmit();
            
            }

            setMessage(content.message);
            setLoading(false);
        }


        if (newPassword !== confirmNewPassword)
            setMessage("Le nouveau mot de passe ne correspond pas");
        else if (props.isAdminRequest)
            apiCallModifyPassword()
        else {
            const messageCheckLogin = await checkLogin(props.userEmail, currentPassword);
            if (messageCheckLogin === "success")
                apiCallModifyPassword()
            else 
                setMessage(messageCheckLogin);
        }
    }

    
    
    return (
        <div>
        <form onSubmit={submit}>
            { !props.isAdminRequest &&
                <>
                <div className="form-group">
                    <label htmlFor="inputCurrentPassword">Mot de passe actuel</label><br />
                    <input type="password" className="form-control" name='inputCurrentPassword' placeholder="Mot de passe actuel" required
                    onChange={e => setCurrentPassword(e.target.value)}
                    />
                </div><br />
                </>
            }
            

            <div className="form-group">
                <label htmlFor="inputNewPassword">Nouveau mot de passe</label><br />
                <input type="password" className="form-control" name='inputNewPassword' placeholder="Nouveau mot de passe" required
                   onChange={e => setNewPassword(e.target.value)}
                />
            </div><br />

            
            <div className="form-group">
                <label htmlFor="inputConfirmNewPassword">Confirmer nouveau mot de passe</label><br />
                <input type="password" className="form-control" name='inputConfirmNewPassword' placeholder="Confirmer nouveau mot de passe" required
                   onChange={e => setConfirmNewPassword(e.target.value)}
                />
            </div><br />

            <button className="w-100 btn btn-lg btn-primary" type="submit">Soumettre</button>
            <p style={{color: 'red'}}>{message}</p>
        </form>
        {loading && <Spinner/>}
        </div>
    );
};

export default ModifyPassword;