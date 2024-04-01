// SignUp.js
import React from 'react';

const SignUp = () => {
    return (
        <div>
            <h2>Sign Up</h2>
            <form>
                <label>
                    Email:
                    <input type="email" name="email" />
                </label>
                <label>
                    Password:
                    <input type="password" name="password" />
                </label>
                <input type="submit" value="Submit" />
            </form>
            <p>Already have an account? <a href="/signin">Sign In</a></p>
        </div>
    );
}

export default SignUp;
