import { useState } from "react";
import { useNavigate } from "react-router-dom";

export const useRegisterHook = () => {
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState<boolean>(false);
  const navigate = useNavigate();

  const register = async (username: string, password: string) => {
    setLoading(true);
    setError(null);
    try {
      const url = "http://localhost:3000/api/register";
      const response = await fetch(url, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          username,
          password,
        }),
      });
      if (!response.ok) {
        throw new Error('Error registering')
      }
      navigate("/login");
    } catch (error) {
      console.error(error);
      setError("Something went wrong please try again");
    } finally {
      setLoading(false);
    }
  };

  return { register, error, loading };
};
