import React, { useState } from 'react';
//import { useNavigate } from 'react-router-dom';
import '../SignUp/SignUp.css';
import SuccessToast from '../Toast/success';
import ErrorToast from '../Toast/err';
const Login = () => {
  // const navigate= useNavigate(); 
  const [formData, setFormData] = useState({
    email: '',
    password: '',
  });

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch('http://localhost:6996/api/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: new URLSearchParams(formData),
      });

      if (response.ok) {
        const result = await response.json();
        SuccessToast(result.message+result.user_id);
        console.log(result);
        //navigate('/dashboard/' + result.user_id);
      } else {

        ErrorToast('Either email/password is incorrect');
        console.error('Failed to submit form');
      }
    } catch (error) {
      ErrorToast(error);
      console.error('Error submitting form:', error);
    }
  };

  return (
    <>
    <main className='main'>
      <form onSubmit={handleSubmit}>
        <label>
          Email:
          <input type="email" name="email" value={formData.email} onChange={handleChange} />
        </label>
        <label>
          Password:
          <input type="password" name="password" value={formData.password} onChange={handleChange} />
        </label>
        <button type="submit">Sign In</button>
      </form>
    </main>
    </>
  );
};

export default Login;
