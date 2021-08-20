import React from 'react'

const Header = (para) => {
  return (
    <div>
      <h1>{para.course}</h1>
    </div>
  )
}

const Content = (paraArr) => {
  let elements = [];
  paraArr.paraArr.forEach(item => {
    elements.push(
      <p key={item.name}>{item.name} {item.content}</p>
    )
  })

  return (
    <div>
      {elements}
    </div>
  )
}

const Total = (paraArr) => {
  let strRet = 'Number of excercises ';

  return (
    <div>
      <p>{strRet} {paraArr.excesers[0] + paraArr.excesers[1] + paraArr.excesers[2]}</p>
    </div>
  )
}

const App = () => {
  const course = 'Half Stack application development'
  const part1 = 'Fundamentals of React'
  const exercises1 = 10
  const part2 = 'Using props to pass data'
  const exercises2 = 7
  const part3 = 'State of a component'
  const exercises3 = 14

  return (
    <div>
      <Header course={course} />
      <Content paraArr={[
        {name: part1, content: exercises1},
        {name: part2, content: exercises2},
        {name: part3, content: exercises3},
      ]} />
      <Total excesers={[exercises1, exercises2, exercises3]}/>
    </div>
  )
}

export default App
