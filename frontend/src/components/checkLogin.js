export async function checkLogin(email, password) {
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
    return contentLogin.message;
};