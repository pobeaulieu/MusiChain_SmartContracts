import {Link} from "react-router-dom";

const Nav = (props: any) => {

    let menu;

    if (props.loggedUser.name === undefined) {
        menu = (
            <ul className="navbar-nav me-auto mb-2 mb-md-0">
                <li className="nav-item active">
                    <Link to="/login"  className="nav-link">Se connecter</Link>
                </li>
            </ul>
        )
    } else {
        menu = (
            <ul className="navbar-nav me-auto mb-2 mb-md-0">
                <li className="nav-item active">
                    <Link to="/login" className="nav-link" onClick={props.onLogout}>Se déconnecter</Link>
                </li>
            </ul>
        )
    }

    return (
         <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4" >
            <div className="container-fluid">
                <Link to="/" className="navbar-brand">Accueil</Link>
                { props.loggedUser.adminRole === 1 && props.loggedUser.blocked === 0? (
                    <>  
                    <Link to="/administration" className="navbar-brand">Administration</Link>
                    <Link to="/clients/business" className="navbar-brand">Clients Affaires</Link>
                    <Link to="/clients/residential" className="navbar-brand">Clients Résidentiels</Link>
                    <Link to="/users" className="navbar-brand">Utilisateurs</Link>
                    </>
                ): props.loggedUser.businessRole === 1 && props.loggedUser.residentialRole === 1 && props.loggedUser.blocked === 0 ? (
                    <> 
                    <Link to="/clients/business" className="navbar-brand">Clients Affaires</Link>
                    <Link to="/clients/residential" className="navbar-brand">Clients Résidentiels</Link>
                    </> 
                ) : props.loggedUser.businessRole === 1  && props.loggedUser.blocked === 0? (
                    <Link to="/clients/business" className="navbar-brand">Clients Affaires</Link>
                ) : props.loggedUser.residentialRole === 1 && props.loggedUser.blocked === 0 ? (
                    <Link to="/clients/residential" className="navbar-brand">Clients Résidentiels</Link>
                ) : null
                }
                <div>
                    {menu}
                </div>
            </div>
        </nav>
    );
};

export default Nav;