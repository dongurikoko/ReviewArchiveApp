import logo from './logo.svg';
import {Route,Routes,BrowserRouter} from 'react-router-dom';
import All from './pages/title/all';
import Create from './pages/title/create';
import Single from './pages/title/single';
import Header from './components/header';
import Footer from './components/footer';
import './App.css';


const App = () => {
  return(
    <BrowserRouter>
    <div className="container">
      <Header />
      <Routes>
        <Route path="/" element={<All />} />
        <Route path="/contents/:id" element={<Single />} />
        <Route path="/create" element={<Create />} />
      </Routes>
      <Footer />
      </div>
    </BrowserRouter>  
  )
}
export default App;
