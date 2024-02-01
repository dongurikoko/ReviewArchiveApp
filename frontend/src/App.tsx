import logo from './logo.svg';
import {Route,Routes,BrowserRouter} from 'react-router-dom';
import All from './pages/title/all';
import Create from './pages/title/create';
import Single from './pages/title/single';
import Header from './components/header';
import Footer from './components/footer';
import Update from "./pages/title/update";
import Delete from "./pages/title/delete";
import './App.css';
import { useState } from 'react';


const App = () => {
  const [searchTerm, setSearchTerm] = useState('');
  return(
    <BrowserRouter>
    <div className="container">
      <Header searchTerm={searchTerm} setSearchTerm={setSearchTerm} />
      <Routes>
        <Route path="/" element={<All searchTerm={searchTerm} setSearchTerm={setSearchTerm}/>} />
        <Route path="/contents/:id" element={<Single />} />
        <Route path="/content/new" element={<Create />} />
        <Route path="/content/update/:id" element={<Update />} />
        <Route path="/content/delete/:id" element={<Delete />} />
      </Routes>
      <Footer />
      </div>
    </BrowserRouter>  
  )
}
export default App;
