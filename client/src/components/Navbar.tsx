
export default function Navbar() {

	return (
		<div className="pt-2">
            <div className="flex justify-between items-center">
                <p  className="text-black font-bold text-[44px] text-center w-full border-b">YourTodo</p>
            </div>
            <div className="flex flex-row gap-3 w-full justify-end pr-2 py-2">
                <a className="bg-slate-400 border border-bg-slate-600 px-3 py-2 rounded-md hover:border-slate-400 text-[16px] cursor-pointer" href="">Login</a>
                <a className="bg-slate-400 border border-bg-slate-600 px-3 py-2 rounded-md hover:border-slate-400 text-[16px] cursor-pointer" href="">Register</a>
            </div>
        </div>
	);
}
