import './UserTable.css'; // <-- Aquí es donde llamamos al CSS

export function UserTable({ users, onEdit, onDelete }) {
  return (
    <table className="mi-tabla">
      <thead>
        <tr>
          <th>ID</th>
          <th>Nombre</th>
          <th>Email</th>
          <th>Rol</th>
          <th>Acciones</th>
        </tr>
      </thead>
      <tbody>
        {users.map((u) => (
          <tr key={u.id}>
            <td>{u.id}</td>
            <td>{u.name} {u.last_name}</td>
            <td>{u.email}</td>
            <td>{u.is_admin ? 'Administrador' : 'Usuario'}</td>
            <td>
              <button className="btn-editar" onClick={() => onEdit(u)}>Editar</button>
              <button className="btn-eliminar" onClick={() => onDelete(u.id)}>Eliminar</button>
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}