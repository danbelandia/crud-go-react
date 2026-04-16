import { useState, useEffect } from 'react';
import { fetchUsers, createUser, updateUser, deleteUser } from './api';
import { UserForm } from './components/UserForm';
import { UserTable } from './components/UserTable';

const initialForm = { id: '', name: '', last_name: '', email: '', is_admin: false };

function App() {
  const [users, setUsers] = useState([]);
  const [page, setPage] = useState(1);
  const [formData, setFormData] = useState(initialForm);

  const loadUsers = async () => {
    try {
      const data = await fetchUsers(page, 5);
      setUsers(data || []); 
    } catch (error) {
      console.error(error);
    }
  };

  useEffect(() => {
    const getInitialData = async () => {
      await loadUsers();
    }
    getInitialData();
  }, [page]);

  const handleChange = (e) => {
    const { name, value, type, checked } = e.target;
    setFormData({
      ...formData,
      [name]: type === 'checkbox' ? checked : value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      if (formData.id) {
        const dataAEnviar = { ...formData, id: Number(formData.id) };
        await updateUser(formData.id, dataAEnviar);
      } else {
        const nuevoUsuario = {
          name: formData.name,
          last_name: formData.last_name,
          email: formData.email,
          is_admin: formData.is_admin
        };
        await createUser(nuevoUsuario);
      }
      setFormData(initialForm);
      loadUsers();
    } catch (error) {
      alert(error.message);
    }
  };

  const handleEdit = (user) => {
    setFormData({
      id: user.id,
      name: user.name,
      last_name: user.last_name,
      email: user.email,
      is_admin: user.is_admin
    });
  };

  const handleDelete = async (id) => {
    if (!window.confirm("¿Seguro que deseas eliminarlo?")) return;
    try {
      await deleteUser(id);
      loadUsers();
    } catch (error) {
      alert(error.message);
    }
  };

  return (
    <div style={{ padding: '20px', fontFamily: 'system-ui', maxWidth: '1000px', margin: '0 auto' }}>
      <h1 class="page-title">Panel de Usuarios</h1>

      <UserForm 
        formData={formData} 
        onChange={handleChange} 
        onSubmit={handleSubmit} 
        onCancel={() => setFormData(initialForm)} 
      />

      <UserTable 
        users={users} 
        onEdit={handleEdit} 
        onDelete={handleDelete} 
      />

<div className="paginacion-contenedor">
  <button 
    className="btn-paginacion"
    onClick={() => setPage(page - 1)} 
    disabled={page === 1}
  >
    Anterior
  </button>
  
  {}
  <div className="paginacion-texto">
    Página <span>{page}</span>
  </div>
  
  <button 
    className="btn-paginacion"
    onClick={() => setPage(page + 1)} 
    disabled={users.length < 5}
  >
    Siguiente
  </button>
</div>
    </div>
  );
}

export default App;