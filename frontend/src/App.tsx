import logo from './logo.svg';
import {Route,Routes,BrowserRouter} from 'react-router-dom';
import Main from './pages/title/main';
import Register from './pages/title/register';
import Contents from './pages/content/contents';
import './App.css';


const App = () => {
  return(
    <BrowserRouter>
      <Routes>
        <Route path="/main" element={<Main />} />
        <Route path="/register" element={<Register />} />
        <Route path="/contents" element={<Contents />} />
      </Routes>
    </BrowserRouter>  
  )
}
export default App;
