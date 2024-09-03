import { useState } from "react";
import { useNavigate } from "react-router-dom";

export const useLogin = () => {
  const [loading, setLoading] = useState<boolean>(false)
  const [error, setError] = useState<string | null>(null)
  const navigate = useNavigate();

  const login = async (username: string, password: string) => {
    setLoading(true)
    setError(null)
    try {
      const url = 'http://localhost:3000/api/login';
      const response = await fetch(url, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          username,
          password
        })
      });
      if (!response.ok) {
        throw new Error('Failed to login')
      }
      const data = await response.json()
      localStorage.setItem('token', data.token)
      navigate('/')
    } catch (error) {
      console.error(error)
      setError('Error logging in')
    } finally {
      setLoading(false)
    }
  };
  return {login, error, loading}
};
