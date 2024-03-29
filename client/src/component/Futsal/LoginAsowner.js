
import React, {useState} from 'react';
import '../SignUp/SignUp.css';
import SuccessToast from '../Toast/success';
import ErrorToast from '../Toast/err';
const LoginAsowner=()=>{
   const [formData, setFormData] = useState({
    email: '',
    password: '',
  });

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };
  const handleSubmit= async (e)=>{

    e.preventDefault();

    try {
      const response = await fetch('http://localhost:6996/loginasowner', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: new URLSearchParams(formData),
      });

      if (response.ok) {
        const result = await response.json();
        console.log(result);
        SuccessToast(result.message);
      } else {
        console.error('Failed to submit form');
        ErrorToast('Futsal is not registered');
      }
    } catch (error) {
      console.error('Error submitting form:', error);
      ErrorToast('Error submitting form'+error);
    }
  }
return(
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
      <button type="submit">Log In</button>
    </form>
    </main>
  );
};

export default LoginAsowner;

