import { useEffect, useState } from "react";
import moment from 'moment';
import Register from "../components/Register";
import ModifyPassword from "../components/ModifyPassword";

interface User {
    id: number;
    name: string;
    email: string;
    adminRole: number;
    businessRole: number;
    residentialRole: number;
    lastModified: string;
    lastLoginAttempt: string;
    loginAttemptCount: number;
    blocked: number;
}

const Users = (props: any) => {
    const [users, setUsers] = useState<User[]>([] as User[]);
    const [loading, setLoading] = useState(true);
    const [visibleUserActions, setVisibleUserActions] = useState<number |null>(null);
    const [isRegisterModalOpen, setIsRegisterModalOpen] = useState(false);
    const [isResetPasswordModalOpen, setIsResetPasswordModalOpen] = useState(false);
    const [resetUserId, setResetUserId] = useState(0 as number);

    useEffect(() => {
        fetchUsers();
    }, []);

    async function fetchUsers() {
        const response = await fetch('https://localhost:8000/api/users', {
            headers: {'Content-Type': 'application/json'},
            credentials: 'include',
        });
        const data = await response.json();
        const sortedUsers = data.sort((a: { id: number; }, b: { id: number; }) => a.id - b.id);
        setUsers(sortedUsers);
        setLoading(false);
    }

    const handleToggleUserDiv = (userId: number) => {
        setVisibleUserActions((prevVisibleUserDiv: number | null) => {
            if (prevVisibleUserDiv === userId) {
                return null;
            } else {
                return userId;
            }
        });
    };

    function convertDate(date: string) {
        const dt = moment("2023-03-25T11:46:47.496262-04:00");
        return dt.format("Do MMMM YYYY, h:mm:ss A z");
    }

    // ROLES MANAGEMENT
    const handleRemoveRole = async (userId: number, role: string) => {
        const roles = newRoles(userId, role, 0);
        callUserRoleAPI(roles);
    }

    const handleAddRole = async (userId: number, role: string) => {
        const roles = newRoles(userId, role, 1);
        callUserRoleAPI(roles);
    }

    function newRoles(userId: number, role: string, value: number) {
        const user = users.find((user) => user.id === userId);
        const currentRoles = {
            userid: user?.id,
            adminRole: user?.adminRole,
            businessRole: user?.businessRole,
            residentialRole: user?.residentialRole,
        };

        if (role === "admin")  currentRoles.adminRole = value;
        if (role === "business") currentRoles.businessRole = value;
        if (role === "residential") currentRoles.residentialRole = value;

        return currentRoles;
    }

    const callUserRoleAPI = async (roles: object) => {
        await fetch('https://localhost:8000/api/userrole', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            credentials: 'include',
            body: JSON.stringify(roles),
        });
        fetchUsers();
    }

    // MODAL MANAGEMENT
    const handleRegisterOpenModal = () => {
        setIsRegisterModalOpen(true);
    };
    
    const handleRegisterCloseModal = () => {
        setIsRegisterModalOpen(false);
        fetchUsers();
    };
      
    function RegisterModal(props: any) {
        return (
            <div className="modal">
            <div className="modal-content">
                <h2>Créer un utilisateur</h2>
                <Register onSubmit={props.onClose}/>
                <button className="btn btn-danger" onClick={props.onClose} style={{marginTop: "20px"}}>Fermer</button>
            </div>
            </div>
        );
    }

    const handleResetPasswordOpenModal = (userId: number) => {
        setResetUserId(userId);
        setIsResetPasswordModalOpen(true);
    };

    const handleResetPasswordCloseModal = () => {
        setIsResetPasswordModalOpen(false);
        fetchUsers();
    };

    function ResetPasswordModal(props: any) {
        const resetUserEmail = users.find((user) => user.id === resetUserId)?.email;

        if  (resetUserEmail !== undefined) {
            return (
                <div className="modal">
                <div className="modal-content">
                    <h2>Configurer un nouveau mot de passe</h2>
                    <ModifyPassword onSubmit={props.onClose} userEmail={resetUserEmail} isAdminRequest={true}/>
                    <button className="btn btn-danger" onClick={props.onClose} style={{marginTop: "20px"}}>Fermer</button>
                </div>
                </div>
            );
        }

        return (
            <div className="modal">
                <div className="modal-content">
                    <h2>Configurer un nouveau mot de passe</h2>
                    <p style={{color: "red"}}>Une erreur est survenue</p>
                    <button onClick={props.onClose} style={{marginTop: "20px"}}>Fermer</button>
                </div>
            </div>
        );
    }

    // RENDER
    if (props.loggedUser.adminRole !== 1 && props.loggedUser.blocked !== 0) {
        return (
            <div>
                Vous n'êtes pas autorisé à accéder à cette page
            </div>
        );
    }

    if (loading) {
        return <div>
            <h2>Utilisateurs</h2>
            Loading...</div>;
    }

    return (
        <div>
            <h2>Utilisateurs</h2>
            <button className="btn btn-success" onClick={handleRegisterOpenModal}> + Ajouter utilisateur</button>
            <br />
            {isRegisterModalOpen && <RegisterModal onClose={handleRegisterCloseModal}/>}
            {isResetPasswordModalOpen && <ResetPasswordModal onClose={handleResetPasswordCloseModal} userId={resetUserId} />}
            <br />
            <table className="type-table">
                <thead>
                <tr>
                    <th>Id</th>
                    <th>Nom d'utilisateur</th>
                    <th>Rôles</th>
                    <th>Bloqué</th>
                    <th>Mot de passe modifié</th>
                    <th>Dernière tentative de connexion échouée</th>
                    <th>Total tentatives échouées</th>
                    <th>Actions</th>
                </tr>
                </thead>
                <tbody>
                {users.map((user) => (
                    <tr key={user.id}>
                        <td>{user.id}</td>
                        <td>{user.name}</td>
                        <td>
                            {user.adminRole === 1 && <div>Admin<br/></div>}
                            {user.businessRole === 1 &&  <span>Business<br/></span>}
                            {user.residentialRole === 1 && <span>Residential<br/></span>}
                        </td>
                        <td>
                            {user.blocked === 1 ? (
                                <span><b>Bloqué</b></span>
                            ) 
                            : user.blocked === 2 ? (
                                <span><b>Attente changement</b></span>
                            ) : null }

                        </td>
                        <td>{convertDate(user.lastModified)}</td>
                        <td>
                            {user.lastLoginAttempt !== "0000-12-31T19:00:00-05:00" && <span>{convertDate(user.lastLoginAttempt)}</span>}
                        </td>
                        <td>{user.loginAttemptCount}</td>
                        <td>
                            <button className=" dropdown-toggle btn btn-dark" onClick={() => handleToggleUserDiv(user.id)}>
                                Actions
                            </button>
                            {visibleUserActions === user.id && (
                                <div className="hidden-div">
                                    {(user.adminRole === 0) ? (
                                        <div><button className="btn btn-success" onClick={() => handleAddRole(user.id, "admin")}>Ajouter rôle admin</button><br /></div>
                                    ) : user.id !== props.loggedUser.id ? (
                                        <div><button className="btn btn-danger" onClick={() => handleRemoveRole(user.id, "admin")}>Retirer rôle admin</button><br /></div>
                                    ) : null
                                    }

                                    {user.businessRole === 0 ? (
                                    <div><button className="btn btn-success" onClick={() => handleAddRole(user.id, "business")}>Ajouter rôle business</button><br /></div>
                                    ) : (
                                        <div><button className="btn btn-danger" onClick={() => handleRemoveRole(user.id, "business")}>Retirer rôle business</button><br /></div>
                                    )}

                                    {user.residentialRole === 0 ? (
                                        <div><button className="btn btn-success" onClick={() => handleAddRole(user.id, "residential")}>Ajouter rôle residential</button><br /></div>
                                    ) : (
                                        <div><button className="btn btn-danger" onClick={() => handleRemoveRole(user.id, "residential")}>Retirer rôle residential</button><br /></div>
                                    )}

                                    {user.blocked === 1 ? (
                                        <button className="btn btn-warning" onClick={() => handleResetPasswordOpenModal(user.id)}>Débloquer et réinitialiser mot de passe</button>
                                    ) : (
                                        <button className="btn btn-warning" onClick={() => handleResetPasswordOpenModal(user.id)}>Réinitialiser mot de passe</button>
                                    )}
                                </div>
                            )}
                        </td>
                    </tr>
                ))}
                </tbody>
            </table>
        </div>
    );
};

export default Users;