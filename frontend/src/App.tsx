import logo from './logo.svg';
import {Route,Routes,BrowserRouter} from 'react-router-dom';
import All from './pages/title/all';
import Create from './pages/title/create';
import Single from './pages/title/single';
import Header from './components/header';
import Footer from './components/footer';
import Update from "./pages/title/update";
import Delete from "./pages/title/delete";
import SignIn from './pages/title/signin';
import SignUp from './pages/title/signup';
import './App.css';
import { useState } from 'react';
import { AuthProvider } from './context/AuthContext';

const App = () => {
  const [searchTerm, setSearchTerm] = useState('');
  return(
    <AuthProvider>
    <BrowserRouter>
    <div className="container">
      <Header searchTerm={searchTerm} setSearchTerm={setSearchTerm} />
      <Routes>
        <Route path="/signup" element={<SignUp />} />
        <Route path="/" element={<All searchTerm={searchTerm} setSearchTerm={setSearchTerm}/>} />
        <Route path="/contents/:id" element={<Single />} />
        <Route path="/content/new" element={<Create />} />
        <Route path="/content/update/:id" element={<Update />} />
        <Route path="/content/delete/:id" element={<Delete />} />
        <Route path="/signin" element={<SignIn />} />
      </Routes>
      <Footer />
      </div>
    </BrowserRouter>  
    </AuthProvider>
  )
}
export default App;
