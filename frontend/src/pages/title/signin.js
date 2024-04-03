// SignIn.js
import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { auth } from '../../main';
import { signInWithEmailAndPassword } from 'firebase/auth';

const SignIn = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    const handleChangeEmail = (event) => {
        setEmail(event.target.value);
    }
    
    const handleChangePassword = (event) => {
        setPassword(event.target.value);
    }
    
    const handleSubmit = async (event) => {
        event.preventDefault(); // フォームのデフォルト送信を防止
        try {
            const userCredential = await signInWithEmailAndPassword(auth, email, password);
            const user = userCredential.user;
            const idToken = await user.getIdToken();
            localStorage.setItem('jwt', idToken.toString());
            console.log(localStorage.getItem('jwt'));
            navigate("/");
        } catch (err) {
            alert(err.message);
            console.error(err);
        }
    };
        
    return (
        <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
            <h1 style={{ marginTop: '10px' }}>Sign In</h1>
            <form onSubmit={handleSubmit}>
                <div>
                    <label>
                        Email:
                        <input 
                            type="email" 
                            name="email" 
                            value={email} // ステートをバインド
                            onChange={handleChangeEmail} // 入力値の変更をステートに反映
                        />
                    </label>
                </div>
                <div>
                    <label>
                        Password:
                        <input 
                            type="password" 
                            name="password" 
                            value={password} // ステートをバインド
                            onChange={handleChangePassword} // 入力値の変更をステートに反映
                        />
                    </label>
                </div>
                <button type="submit" className="register-button">Submit</button>
            </form>
            <p>Don't have an account? <a href="/signup">Sign Up</a></p>
        </div>
    );
}

export default SignIn;
