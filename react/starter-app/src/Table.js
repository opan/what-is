import React, { Component } from "react"

// class Table extends Component {
//   render() {
//     return (
//       <table>
//         <thead>
//           <tr>
//             <th>Name</th>
//             <th>Job</th>
//           </tr>
//         </thead>
//         <tbody>
//           <tr>
//             <td>Charlie</td>
//             <td>Janitor</td>
//           </tr>
//           <tr>
//             <td>Mac</td>
//             <td>Bouncer</td>
//           </tr>
//           <tr>
//             <td>Dee</td>
//             <td>Aspiring actress</td>
//           </tr>
//           <tr>
//             <td>Dennis</td>
//             <td>Bartender</td>
//           </tr>
//         </tbody>
//       </table>
//     )
//   }
// }


const TableHeader = () => {
  return (
    <thead>
      <tr>
        <th>Name</th>
        <th>Job</th>
      </tr>
    </thead>
  )
}

const TableBody = (props) => {
  const rows = props.charactersData.map((row, index) => {
    return (
      <tr key={index}>
        <td>{row.name}</td>
        <td>{row.job}</td>
        <td>
          <button onClick={() => props.removeCharacter(index) }>Delete</button>
        </td>
      </tr>
    )
  })
  return (
    <tbody>{rows}</tbody>
  )

}

const Table = (props) => {
  const { charactersData, removeCharacter } = props

  return (
    <table>
      <TableHeader />
      <TableBody charactersData={charactersData} removeCharacter={removeCharacter} />
    </table>
  )
}

export default Table
