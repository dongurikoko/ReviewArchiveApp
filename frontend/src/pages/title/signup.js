// SignUp.js
import React from 'react';

const SignUp = () => {
    return (
        <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
            <h1 style={{ marginTop: '10px' }}>Sign Up</h1>
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
            <p>Already have an account? <a href="/signin">Sign In</a></p>
        </div>
    );
}

export default SignUp;
