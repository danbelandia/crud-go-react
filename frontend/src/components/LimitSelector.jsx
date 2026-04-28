export function LimitSelector({ limit, onChange }) {
  const options = [5, 10, 20, 50, 100]; // Opciones de límite en el filtro disponibles

  return (
    <div style={{ marginBottom: '15px' }}>
      <label style={{ marginRight: '10px', fontWeight: 'bold' }}>Mostrar:</label>
      <select 
        value={limit} 
        onChange={(e) => onChange(Number(e.target.value))}
        style={{ padding: '5px', borderRadius: '4px', border: '1px solid #BFC9D1' }}
      >
        {options.map(opt => (
          <option key={opt} value={opt}>{opt} usuarios</option>
        ))}
      </select>
    </div>
  );
}