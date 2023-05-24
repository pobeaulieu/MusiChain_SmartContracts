import ModifyPassword from "../components/ModifyPassword";

const Home = (props: any) => {

    return (
        <div>
            {props.loggedUser.blocked === 2 && <h2 style={{color: "red"}}>Vous devez changer votre mot de passe pour accéder au site</h2>}
            {props.loggedUser.name ? (
                <div>
                    <h2>Bonjour {props.loggedUser.name}</h2>
                    <br />
                    <h3>Changer votre mot de passe</h3>
                    <ModifyPassword onSubmit={() => props.loadUser } userEmail={props.loggedUser.email} isAdminRequest={false}/>
                </div>
            ) : <h2>Vous n'êtes pas connecté</h2>
            }
        </div>
    );
};



export default Home;