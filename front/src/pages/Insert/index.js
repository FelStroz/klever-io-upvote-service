import React, { useReducer, useState } from 'react';
import { Link } from 'react-router-dom';
import TextField from '@material-ui/core/TextField';
import { makeStyles } from '@material-ui/core/styles';
import './index.css'

const useStyles = makeStyles((theme) => ({
    root: {
        '& .MuiTextField-root': {
            margin: theme.spacing(1),
            width: 300,
        },
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center'
    },
}));

const formReducer = (state, event) => {
    return {
        ...state,
        name: event.value
    }
}

function InsertForm() {
    const classes = useStyles();
    const [erro, setErro] = useState(false);
    const [erroMessage, setErroMessage] = useState("");
    const [formData, setFormData] = useReducer(formReducer, {});

    function handleSubmit() {
        if (formData.name == "" || formData.name == undefined) {
            setErro(true);
            setErroMessage("Name must not be empty")
            return
        }

        fetch(`http://localhost:8000/crypto?name=${formData.name}`, {
            "method": "POST",
            "headers": {}
        })
            .then(async response => {
                let res = await response.json();
                if (res.erro) {
                    setErro(true);
                    setErroMessage("Crypto already exists")
                } else {
                    setErro(false);
                    setErroMessage("")
                    window.location.replace('http://localhost:3000')
                }
            })
            .catch(err => {
                setErro(true);
                setErroMessage("Crypto already exists")
            });
    }

    const handleChange = event => {
        setFormData({
            value: event.target.value
        });
    }

    return (
        <div className="container">
            <h1>Insert Cryptocurrency</h1>

            <TextField
                error={erro}
                id="outlined-error-helper-text"
                label="Crypto Name"
                defaultValue=""
                helperText={erroMessage}
                variant="outlined"
                onChange={handleChange}
            />

            <div className="buttonInsert" onClick={() => handleSubmit()}><text className="buttonText">Insert</text></div>
        </div>
    );
}

export default InsertForm;