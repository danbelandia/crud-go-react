const BASE_URL = 'http://localhost:8080/usuarios';


const getHeaders = () => ({
  'Content-Type': 'application/json',
  'X-Role': 'super-clave-admin-123'
});

//GET Paginado
export const fetchUsers = async (page = 1, limit = 5, search = '') => {
  const querySearch = encodeURIComponent(search); //Codificamos para evitar problemas con caracteres especiales
  const response = await fetch(`${BASE_URL}?page=${page}&limit=${limit}&search=${querySearch}`); //Obtiene la pagina actual y el límite dinámico, además del término de búsqueda
  if (!response.ok) throw new Error('Error al obtener usuarios.');
  return response.json();
};

//POST
export const createUser = async (userData) => {
  const response = await fetch(BASE_URL, {
    method: 'POST',
    headers: getHeaders(),
    body: JSON.stringify(userData),
  });
  if (!response.ok) throw new Error('Error al crear usuario.');
  return response.json();
};

// PUT
export const updateUser = async (id, userData) => {
  const response = await fetch(`${BASE_URL}/${id}`, {
    method: 'PUT',
    headers: getHeaders(),
    body: JSON.stringify(userData),
  });
  if (!response.ok) throw new Error('Error al actualizar usuario.');
  return response.json();
};

// DELETE
export const deleteUser = async (id) => {
  const response = await fetch(`${BASE_URL}/${id}`, {
    method: 'DELETE',
    headers: getHeaders(),
  });
  if (!response.ok) throw new Error('Error al eliminar usuario.');
  return response.json();
};