const BASE_URL = 'http://localhost:8080/usuarios';

// Centralizamos los headers para no repetirlos
const getHeaders = () => ({
  'Content-Type': 'application/json',
  'X-Role': 'super-clave-admin-123'
});

//GET Paginado
export const fetchUsers = async (page = 1, limit = 5) => {
  const response = await fetch(`${BASE_URL}?page=${page}&limit=${limit}`);
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