import { useEffect, useState } from 'react';
import { Spinner, Tab, Tabs } from 'react-bootstrap';
import { checkLogin } from '../components/checkLogin';

interface LoginForm {
    id: number;
    maxLoginAttemptCount: number;
    loginTimeInterval: number;
    twofa:boolean;
}

interface PasswordForm {
    Id: number;
    minLength: number;
    requireNumber: boolean;
    requireLower: boolean;
    requireUpper: boolean;
    requireSymbol: boolean;
    historySize: number;
}

const Administration = (props: any) => {
    const [loginFormData, setLoginFormData] = useState<LoginForm>({
        id: 0,
        maxLoginAttemptCount: 0,
        loginTimeInterval: 0
    } as LoginForm);
    const [passwordFormData, setPasswordFormData] = useState<PasswordForm>({
        Id: 0,
        minLength: 0,
        requireNumber: false,
        requireLower: false,
        requireUpper: false,
        requireSymbol: false,
        historySize: 0
    } as PasswordForm);
    const [loginPolicyPassword, setLoginPolicyPassword] = useState('');
    const [loginPolicyMessage, setLoginPolicyMessage] = useState('');
    const [passwordPolicyPassword, setPasswordPolicyPassword] = useState('');
    const [passwordPolicyMessage, setPasswordPolicyMessage] = useState('');
    const [loading, setLoading] = useState(false);


    useEffect(() => {
        async function fetchLoginPolicy() {
            const response = await fetch('https://localhost:8000/api/getloginpolicy', {
                headers: {'Content-Type': 'application/json'},
                credentials: 'include',
            });
            const data = await response.json();
            setLoginFormData(data);
        }

        async function fetchPasswordPolicy() {
            const response = await fetch('https://localhost:8000/api/getpasswordpolicy', {
                headers: {'Content-Type': 'application/json'},
                credentials: 'include',
            });
            const data = await response.json();
            setPasswordFormData(data);
        }

        fetchLoginPolicy();
        fetchPasswordPolicy();
    }, []);  

    // LOGIN POLICY
    const handleLoginSubmit = async (event: any) => {
        setLoading(true);
        event.preventDefault();
        
        const messageCheckLogin = await checkLogin(props.loggedUser.email, loginPolicyPassword);
        if (messageCheckLogin === "success"){
            updateLoginPolicy();
        }
        else{
            setLoginPolicyMessage(messageCheckLogin);
        }
        setLoading(false);
    };

    const handleLoginPolicyChange = (event: any) => {
        const { name, value } = event.target;
        setLoginFormData((prevValues) => ({
            ...prevValues,
            [name]: parseInt(value)
        }));
    };

    const updateLoginPolicy = async () => {
        const maxAttempt = {
            max: loginFormData.maxLoginAttemptCount
        }

        const timeInterval = {
            seconds: loginFormData.loginTimeInterval
        }

        await fetch('https://localhost:8000/api/maxloginattempts', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            credentials: 'include',
            body: JSON.stringify(maxAttempt),
        });

        await fetch('https://localhost:8000/api/logintimeinterval', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            credentials: 'include',
            body: JSON.stringify(timeInterval),
        });
    }

    // PASSWORD POLICY
    const handlePasswordSubmit = async (event: any) => {
        setLoading(true);
        event.preventDefault();

        const messageCheckLogin = await checkLogin(props.loggedUser.email, passwordPolicyPassword);
        if (messageCheckLogin === "success"){
            updatePasswordPolicy();
        }
        else{
            setPasswordPolicyMessage(messageCheckLogin);
        }
 
        setLoading(false);

        
    };

    const handlePasswordPolicyChange = (event: any) => {
        const { name, value } = event.target;
        if (name === "requireNumber" || name === "requireLower" || name === "requireUpper" || name === "requireSymbol") {
            setPasswordFormData((prevValues) => ({
                ...prevValues,
                [name]: event.target.checked
            }));
        }
        else {
            setPasswordFormData((prevValues) => ({
                ...prevValues,
                [name]: parseInt(value)
            }));
        }
    };

    const updatePasswordPolicy = async () => {
        const data = {
            minLength: passwordFormData.minLength,
            requireNumber: passwordFormData.requireNumber,
            requireLower: passwordFormData.requireLower,
            requireUpper: passwordFormData.requireUpper,
            requireSymbol: passwordFormData.requireSymbol,
            historySize: passwordFormData.historySize
        }

        await fetch('https://localhost:8000/api/passwordpolicy', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            credentials: 'include',
            body: JSON.stringify(data),
        });
    }


    // RENDERING
    if (props.loggedUser.adminRole !== 1 && props.loggedUser.blocked !== 0) {
        return (
            <div>
                Vous n'êtes pas autorisé à accéder à cette page
            </div>
        );
    }

    return (
        <div>
            <h2>{'Administration'}</h2>
            <Tabs defaultActiveKey="connection" className="mb-3" >
                <Tab eventKey="connection" title="Politique de connexion">
                    <form onSubmit={handleLoginSubmit}>
                        <div className="form-group">
                            <label htmlFor="inputMaxLoginAttemptCount">Nombre de connexion</label><br />
                            <input type="number" id="inputMaxLoginAttemptCount" name="maxLoginAttemptCount" onChange={handleLoginPolicyChange} value={loginFormData.maxLoginAttemptCount} />
                        </div><br />

                        <div className="form-group">
                            <label htmlFor="inputLoginTimeInterval">Intervalle de temps</label><br />
                            <input type="number" id="inputLoginTimeInterval" name="loginTimeInterval" onChange={handleLoginPolicyChange} value={loginFormData.loginTimeInterval} />
                        </div><br />

                        <div className="form-group">
                            <label htmlFor="inputPassword">Entrez votre mot de passe pour sauvegarder les changements</label><br />
                            <input type="password" id='inputPassword' name='password' placeholder="Mot de passe actuel" required
                                onChange={e => setLoginPolicyPassword(e.target.value)}
                            />
                        </div><br />

                        <button type="submit" className="btn btn-primary">Enregistrer</button>
                        <p style={{color: 'red'}}>{loginPolicyMessage}</p>
                    </form>
                </Tab>
                <Tab eventKey="complexPxd" title="Politique de mot de passe">
                    <form onSubmit={handlePasswordSubmit}>
                        <div className="form-group">
                            <label htmlFor="inputMinLength">Longueur minimale</label><br />
                            <input type="number" id="inputMinLength" name="minLength" onChange={handlePasswordPolicyChange} value={passwordFormData.minLength} />
                        </div><br />

                        <div className="form-group">
                            <input  style={{ marginRight: '8px' }} type="checkbox" id="inputRequireNumber" name="requireNumber" onChange={handlePasswordPolicyChange} checked={passwordFormData.requireNumber} />
                            <label htmlFor="inputRequireNumber">Requiert un chiffre</label>
                        </div><br />

                        <div className="form-group">
                            <input style={{ marginRight: '8px' }} type="checkbox" id="inputRequireLower" name="requireLower" onChange={handlePasswordPolicyChange} checked={passwordFormData.requireLower} />
                            <label htmlFor="inputRequireLower">Requiert une lettre minuscule</label>
                        </div><br />

                        <div className="form-group">
                            <input style={{ marginRight: '8px' }} type="checkbox" id="inputRequireUpper" name="requireUpper" onChange={handlePasswordPolicyChange} checked={passwordFormData.requireUpper} />
                            <label htmlFor="inputRequireUpper">Requiert une lettre majuscule</label>
                        </div><br />

                        <div className="form-group">
                            <input style={{ marginRight: '8px' }} type="checkbox" id="inputRequireSymbol" name="requireSymbol" onChange={handlePasswordPolicyChange} checked={passwordFormData.requireSymbol} />
                            <label htmlFor="inputRequireSymbol">Requiert un symbole</label>
                        </div><br />

                        <div className="form-group">
                            <label htmlFor="inputHistory">Historique</label><br />
                            <input type="number" id="inputHistory" name="historySize" onChange={handlePasswordPolicyChange} value={passwordFormData.historySize} />
                        </div><br />

                        <div className="form-group">
                            <label htmlFor="inputPassword">Entrez votre mot de passe pour sauvegarder les changements</label><br />
                            <input type="password" id='inputPassword' name='password' placeholder="Mot de passe actuel" required
                                onChange={e => setPasswordPolicyPassword(e.target.value)}
                            />
                        </div><br />

                        <button type="submit" className="btn btn-primary">Enregistrer</button>
                        <p style={{color: 'red'}}>{passwordPolicyMessage}</p>
                        
                    </form>
                </Tab>
            </Tabs>
            {loading && <Spinner/>}
           
        </div>
        
    );
};

export default Administration;