// SignIn.js
import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { auth } from '../../main';
import { signInWithEmailAndPassword } from 'firebase/auth';

const SignIn = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    // フォーム送信時の処理
    const handleSubmit = async (event) => {
        event.preventDefault(); // フォームのデフォルト送信を阻止
        try {
            // Firebaseのサインイン機能を呼び出し
            const userCredential = await signInWithEmailAndPassword(auth, email, password);
            console.log(userCredential); // 成功した場合、ユーザー情報をコンソールに出力
            // サインイン後の処理
            navigate('/all');
        } catch (error) {
            console.error(error); // エラーが発生した場合、コンソールに出力
        }
    }
        

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
                <button type="submit" class="register-button">Submit</button>
            </form>
            <p>Don't have an account? <a href="/">Sign Up</a></p>
        </div>
    );
}

export default SignIn;
