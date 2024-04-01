// SignUp.js
import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom'; 
import { auth } from '../../main';
import { createUserWithEmailAndPassword } from "firebase/auth";

const SignUp = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    // フォーム送信時の処理
    const handleSubmit = async (event) => {
        event.preventDefault(); // フォームのデフォルト送信を阻止
        try {
            // Firebaseのサインアップ機能を呼び出し
            const userCredential = await createUserWithEmailAndPassword(auth, email, password);
            console.log(userCredential); // 成功した場合、ユーザー情報をコンソールに出力
            // サインアップ後の処理
            navigate('/signin');
        } catch (error) {
            console.error(error); // エラーが発生した場合、コンソールに出力
        }
    };

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
                            onChange={(e) => setEmail(e.target.value)} // 入力値の変更をステートに反映
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
                            onChange={(e) => setPassword(e.target.value)} // 入力値の変更をステートに反映
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
