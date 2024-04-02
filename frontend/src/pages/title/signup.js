// SignUp.js
import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom'; 
import { auth } from '../../main';
import { createUserWithEmailAndPassword } from "firebase/auth";

const SignUp = () => {
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
        event.preventDefault(); // デフォルトのイベントをキャンセル
        try {
            await createUserWithEmailAndPassword(auth, email, password);
            navigate("/signin"); // サインインページへリダイレクト
        } catch (err) {
            alert(err.message); // エラーメッセージをアラート表示
            console.error(err); // コンソールにエラー出力
        }
    }

    return (
        <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
            <h1 style={{ marginTop: '10px' }}>Sign Up</h1>
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
                <button type="submit" class="register-button">Submit</button>
            </form>
            <p>Already have an account? <a href="/signin">Sign In</a></p>
        </div>
    );
}

export default SignUp;
