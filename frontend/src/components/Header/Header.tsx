import React from 'react'
import "./Header.css"
import { Button } from '@mui/material'


interface HeaderProps {
    headerName: string;
    number: number;
    setValue: React.Dispatch<React.SetStateAction<number>>;
    value: number;
}

const Header: React.FC<HeaderProps> = ({ headerName, number, setValue, value }) => {

    return (
        <div className='header-container'>
            <div>Header</div>
            <div className="header-name">
                {headerName}
            </div>
            <div className="number">
                {number}
            </div>
            <Button onClick={()=>{setValue(value + 1)}} variant='contained' color='success'>
                count
            </Button>
        </div>
    )
}

export default Header

