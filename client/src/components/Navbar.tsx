import {Link} from 'react-router-dom'
export default function Navbar() {
	return (
		<div className="pt-2">
            <div className="flex justify-between items-center">
                <p  className="text-black font-bold text-[44px] text-center w-full border-b"><Link to="/">YourTodo</Link></p>
            </div>

        </div>
	);
}
