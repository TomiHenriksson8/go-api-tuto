import { Link } from 'react-router-dom';

const Welcome = () => {
  return (
    <div
      className="flex flex-col items-center justify-center h-screen bg-gray-50"
      style={{
        backgroundImage: "url('./outobg.png')",
        backgroundSize: 'cover',
        backgroundPosition: 'center',
        opacity: 0.9
      }}
    >
      <div className="bg-gray-50 px-20 py-16 rounded-lg shadow-lg text-center z-10 border border-gray-300">
        <h1 className="text-4xl font-bold text-gray-800 mb-4">Welcome to YourTodo</h1>
        <p className="text-lg text-black mb-8">Your personal todo manager. Please log in or register to get started.</p>
        <div className="flex justify-center gap-4">
          <Link
            to="/login"
            className="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-700"
          >
            Login
          </Link>
          <Link
            to="/register"
            className="bg-green-500 text-white px-4 py-2 rounded-md hover:bg-green-700"
          >
            Register
          </Link>
        </div>
      </div>
    </div>
  );
};

export default Welcome;
