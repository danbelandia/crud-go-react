export function UserForm({ formData, onChange, onSubmit, onCancel }) {
  return (
    <form className="formulario-caja" onSubmit={onSubmit}>
      <h3>{formData.id ? 'Editar Usuario' : 'Crear Nuevo Usuario'}</h3>
      
      <input 
        type="text" name="name" placeholder="Nombre" required
        value={formData.name} onChange={onChange} style={{ marginRight: '10px' }}
      />
      <input 
        type="text" name="last_name" placeholder="Apellido" required
        value={formData.last_name} onChange={onChange} style={{ marginRight: '10px' }}
      />
      <input 
        type="email" name="email" placeholder="Correo" required
        value={formData.email} onChange={onChange} style={{ marginRight: '10px' }}
      />
      <label style={{ marginRight: '10px' }}>
        <input 
          type="checkbox" name="is_admin" 
          checked={formData.is_admin} onChange={onChange} 
        /> Admin
      </label>
      
      <button type="submit" style={{ padding: '5px 15px', cursor: 'pointer' }}>
        {formData.id ? 'Actualizar' : 'Guardar'}
      </button>
      
      {formData.id && (
        <button type="button" onClick={onCancel} style={{ marginLeft: '10px', padding: '5px 15px', cursor: 'pointer' }}>
          Cancelar
        </button>
      )}
    </form>
  );
}