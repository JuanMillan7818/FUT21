import  { BrowserRouter, Route, Routes } from 'react-router-dom';
import { AppFut21 } from '../components/AppFut21';
import { Prueba } from '../components/Prueba';

export const AppRouter = () => {
    return (
        <BrowserRouter>
        <AppFut21 />
        <Routes>
            <Route path="/" element={<Prueba />}/>
            <Route path="/search" element={<Prueba />}/>
            <Route path="/team" element={<Prueba />}/>
        </Routes>
        </BrowserRouter>
    )
}
