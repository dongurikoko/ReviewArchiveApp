// SignIn.js
import React from 'react';

const SignIn = () => {
    return (
        <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
            <h1 style={{ marginTop: '10px' }}>Sign In</h1>
            <form>
                <div>
                    <label>
                        Email:
                        <input type="email" name="email" />
                    </label>
                </div>
                <div>
                    <label>
                        Password:
                        <input type="password" name="password" />
                    </label>
                </div>
                <input type="submit" value="Submit" />
            </form>
            <p>Don't have an account? <a href="/signup">Sign Up</a></p>
        </div>
    );
}

export default SignIn;
