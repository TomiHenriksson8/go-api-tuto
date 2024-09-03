export const logout = async () => {
  localStorage.removeItem("token");
  window.location.reload();
};
